# Copyright 2019 Google LLC All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
ARG BASE_IMAGE=agones-build-sdk-base:latest
FROM $BASE_IMAGE

# Add installation step required to generate your SDK.
# This should contains tooling required to generate GRPC code, build (if distribution is required) 
# and test your integration.

# code generation scripts
COPY *.sh /root/
RUN chmod +x /root/*.sh