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
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1Agent = factory(root.PolyaxonSdk.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The V1Agent model module.
   * @module model/V1Agent
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>V1Agent</code>.
   * @alias module:model/V1Agent
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1Agent</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1Agent} obj Optional instance to populate.
   * @return {module:model/V1Agent} The populated <code>V1Agent</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('uuid'))
        obj.uuid = ApiClient.convertToType(data['uuid'], 'String');
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('description'))
        obj.description = ApiClient.convertToType(data['description'], 'String');
      if (data.hasOwnProperty('tags'))
        obj.tags = ApiClient.convertToType(data['tags'], ['String']);
      if (data.hasOwnProperty('disabled'))
        obj.disabled = ApiClient.convertToType(data['disabled'], 'Boolean');
      if (data.hasOwnProperty('deleted'))
        obj.deleted = ApiClient.convertToType(data['deleted'], 'Boolean');
      if (data.hasOwnProperty('namespace'))
        obj.namespace = ApiClient.convertToType(data['namespace'], 'String');
      if (data.hasOwnProperty('version_api'))
        obj.version_api = ApiClient.convertToType(data['version_api'], Object);
      if (data.hasOwnProperty('created_at'))
        obj.created_at = ApiClient.convertToType(data['created_at'], 'Date');
      if (data.hasOwnProperty('updated_at'))
        obj.updated_at = ApiClient.convertToType(data['updated_at'], 'Date');
    }
    return obj;
  }

  /**
   * @member {String} uuid
   */
  exports.prototype.uuid = undefined;

  /**
   * @member {String} name
   */
  exports.prototype.name = undefined;

  /**
   * @member {String} description
   */
  exports.prototype.description = undefined;

  /**
   * @member {Array.<String>} tags
   */
  exports.prototype.tags = undefined;

  /**
   * @member {Boolean} disabled
   */
  exports.prototype.disabled = undefined;

  /**
   * @member {Boolean} deleted
   */
  exports.prototype.deleted = undefined;

  /**
   * @member {String} namespace
   */
  exports.prototype.namespace = undefined;

  /**
   * @member {Object} version_api
   */
  exports.prototype.version_api = undefined;

  /**
   * @member {Date} created_at
   */
  exports.prototype.created_at = undefined;

  /**
   * @member {Date} updated_at
   */
  exports.prototype.updated_at = undefined;

  return exports;

}));
