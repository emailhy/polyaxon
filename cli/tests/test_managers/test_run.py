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

from unittest import TestCase

import pytest

from polyaxon_sdk import V1Run

from polyaxon.managers.run import RunManager


@pytest.mark.managers_mark
class TestRunManager(TestCase):
    def test_default_props(self):
        assert RunManager.IS_GLOBAL is False
        assert RunManager.IS_POLYAXON_DIR is True
        assert RunManager.CONFIG_FILE_NAME == ".polyaxonrun"
        assert RunManager.CONFIG == V1Run
