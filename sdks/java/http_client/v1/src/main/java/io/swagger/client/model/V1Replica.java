// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import io.swagger.client.model.V1Container;
import io.swagger.client.model.V1Environment;
import io.swagger.client.model.V1Init;
import io.swagger.client.model.V1Mounts;
import io.swagger.client.model.V1Termination;
import java.io.IOException;

/**
 * V1Replica
 */

public class V1Replica {
  @SerializedName("replicas")
  private Integer replicas = null;

  @SerializedName("environment")
  private V1Environment environment = null;

  @SerializedName("termination")
  private V1Termination termination = null;

  @SerializedName("init")
  private V1Init init = null;

  @SerializedName("mounts")
  private V1Mounts mounts = null;

  @SerializedName("container")
  private V1Container container = null;

  public V1Replica replicas(Integer replicas) {
    this.replicas = replicas;
    return this;
  }

   /**
   * Get replicas
   * @return replicas
  **/
  @ApiModelProperty(value = "")
  public Integer getReplicas() {
    return replicas;
  }

  public void setReplicas(Integer replicas) {
    this.replicas = replicas;
  }

  public V1Replica environment(V1Environment environment) {
    this.environment = environment;
    return this;
  }

   /**
   * Get environment
   * @return environment
  **/
  @ApiModelProperty(value = "")
  public V1Environment getEnvironment() {
    return environment;
  }

  public void setEnvironment(V1Environment environment) {
    this.environment = environment;
  }

  public V1Replica termination(V1Termination termination) {
    this.termination = termination;
    return this;
  }

   /**
   * Get termination
   * @return termination
  **/
  @ApiModelProperty(value = "")
  public V1Termination getTermination() {
    return termination;
  }

  public void setTermination(V1Termination termination) {
    this.termination = termination;
  }

  public V1Replica init(V1Init init) {
    this.init = init;
    return this;
  }

   /**
   * Get init
   * @return init
  **/
  @ApiModelProperty(value = "")
  public V1Init getInit() {
    return init;
  }

  public void setInit(V1Init init) {
    this.init = init;
  }

  public V1Replica mounts(V1Mounts mounts) {
    this.mounts = mounts;
    return this;
  }

   /**
   * Get mounts
   * @return mounts
  **/
  @ApiModelProperty(value = "")
  public V1Mounts getMounts() {
    return mounts;
  }

  public void setMounts(V1Mounts mounts) {
    this.mounts = mounts;
  }

  public V1Replica container(V1Container container) {
    this.container = container;
    return this;
  }

   /**
   * Get container
   * @return container
  **/
  @ApiModelProperty(value = "")
  public V1Container getContainer() {
    return container;
  }

  public void setContainer(V1Container container) {
    this.container = container;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Replica v1Replica = (V1Replica) o;
    return Objects.equals(this.replicas, v1Replica.replicas) &&
        Objects.equals(this.environment, v1Replica.environment) &&
        Objects.equals(this.termination, v1Replica.termination) &&
        Objects.equals(this.init, v1Replica.init) &&
        Objects.equals(this.mounts, v1Replica.mounts) &&
        Objects.equals(this.container, v1Replica.container);
  }

  @Override
  public int hashCode() {
    return Objects.hash(replicas, environment, termination, init, mounts, container);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Replica {\n");
    
    sb.append("    replicas: ").append(toIndentedString(replicas)).append("\n");
    sb.append("    environment: ").append(toIndentedString(environment)).append("\n");
    sb.append("    termination: ").append(toIndentedString(termination)).append("\n");
    sb.append("    init: ").append(toIndentedString(init)).append("\n");
    sb.append("    mounts: ").append(toIndentedString(mounts)).append("\n");
    sb.append("    container: ").append(toIndentedString(container)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

