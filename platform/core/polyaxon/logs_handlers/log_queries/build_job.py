from typing import Iterable

import conf

from logs_handlers.log_queries import base
from logs_handlers.utils import safe_log_job
from options.registry.container_names import CONTAINER_NAME_BUILD_JOBS
from options.registry.k8s import K8S_NAMESPACE
from polyaxon_k8s.manager import K8SManager


def stream_logs(build: 'BuildJob') -> Iterable[str]:
    k8s_manager = K8SManager(namespace=conf.get(K8S_NAMESPACE), in_cluster=True)
    return base.stream_logs(k8s_manager=k8s_manager,
                            pod_id=build.pod_id,
                            container_job_name=conf.get(CONTAINER_NAME_BUILD_JOBS))


def process_logs(build: 'BuildJob', temp: bool = True) -> None:
    k8s_manager = K8SManager(namespace=conf.get(K8S_NAMESPACE), in_cluster=True)
    log_lines = base.process_logs(k8s_manager=k8s_manager,
                                  pod_id=build.pod_id,
                                  container_job_name=conf.get(CONTAINER_NAME_BUILD_JOBS))

    safe_log_job(job_name=build.unique_name, log_lines=log_lines, temp=temp, append=False)
