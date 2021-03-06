#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    OpenAPI spec version: 1.0.0
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import unittest

import polyaxon_sdk
from polyaxon_sdk.api.agents_v1_api import AgentsV1Api  # noqa: E501
from polyaxon_sdk.rest import ApiException


class TestAgentsV1Api(unittest.TestCase):
    """AgentsV1Api unit test stubs"""

    def setUp(self):
        self.api = polyaxon_sdk.api.agents_v1_api.AgentsV1Api()  # noqa: E501

    def tearDown(self):
        pass

    def test_create_agent(self):
        """Test case for create_agent

        List runs  # noqa: E501
        """
        pass

    def test_delete_agent(self):
        """Test case for delete_agent

        Patch run  # noqa: E501
        """
        pass

    def test_get_agent(self):
        """Test case for get_agent

        Create new run  # noqa: E501
        """
        pass

    def test_list_agent_names(self):
        """Test case for list_agent_names

        List bookmarked runs for user  # noqa: E501
        """
        pass

    def test_list_agents(self):
        """Test case for list_agents

        List archived runs for user  # noqa: E501
        """
        pass

    def test_patch_agent(self):
        """Test case for patch_agent

        Update run  # noqa: E501
        """
        pass

    def test_update_agent(self):
        """Test case for update_agent

        Get run  # noqa: E501
        """
        pass


if __name__ == "__main__":
    unittest.main()
