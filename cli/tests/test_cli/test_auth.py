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
from __future__ import absolute_import, division, print_function

import pytest

from mock import patch
from tests.test_cli.utils import BaseCommandTestCase

from polyaxon.cli.auth import logout, whoami


@pytest.mark.cli_mark
class TestCliAuth(BaseCommandTestCase):
    @patch("polyaxon.managers.auth.AuthConfigManager.purge")
    def test_logout(self, get_user):
        self.runner.invoke(logout)
        assert get_user.call_count == 1

    @patch("polyaxon_sdk.UsersV1Api.get_user")
    def test_whoami(self, get_user):
        self.runner.invoke(whoami)
        assert get_user.call_count == 1
