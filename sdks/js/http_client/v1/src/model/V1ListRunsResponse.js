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

/**
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.7
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1Run'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V1Run'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1ListRunsResponse = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1Run);
  }
}(this, function(ApiClient, V1Run) {
  'use strict';




  /**
   * The V1ListRunsResponse model module.
   * @module model/V1ListRunsResponse
   * @version 1.14.4
   */

  /**
   * Constructs a new <code>V1ListRunsResponse</code>.
   * @alias module:model/V1ListRunsResponse
   * @class
   */
  var exports = function() {
    var _this = this;





  };

  /**
   * Constructs a <code>V1ListRunsResponse</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1ListRunsResponse} obj Optional instance to populate.
   * @return {module:model/V1ListRunsResponse} The populated <code>V1ListRunsResponse</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('count')) {
        obj['count'] = ApiClient.convertToType(data['count'], 'Number');
      }
      if (data.hasOwnProperty('results')) {
        obj['results'] = ApiClient.convertToType(data['results'], [V1Run]);
      }
      if (data.hasOwnProperty('previous')) {
        obj['previous'] = ApiClient.convertToType(data['previous'], 'String');
      }
      if (data.hasOwnProperty('next')) {
        obj['next'] = ApiClient.convertToType(data['next'], 'String');
      }
    }
    return obj;
  }

  /**
   * @member {Number} count
   */
  exports.prototype['count'] = undefined;
  /**
   * @member {Array.<module:model/V1Run>} results
   */
  exports.prototype['results'] = undefined;
  /**
   * @member {String} previous
   */
  exports.prototype['previous'] = undefined;
  /**
   * @member {String} next
   */
  exports.prototype['next'] = undefined;



  return exports;
}));


