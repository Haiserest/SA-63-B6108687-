/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntStatusTrackingEdges,
    EntStatusTrackingEdgesFromJSON,
    EntStatusTrackingEdgesFromJSONTyped,
    EntStatusTrackingEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatusTracking
 */
export interface EntStatusTracking {
    /**
     * 
     * @type {EntStatusTrackingEdges}
     * @memberof EntStatusTracking
     */
    edges?: EntStatusTrackingEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntStatusTracking
     */
    id?: number;
    /**
     * Status holds the value of the "status" field.
     * @type {string}
     * @memberof EntStatusTracking
     */
    status?: string;
}

export function EntStatusTrackingFromJSON(json: any): EntStatusTracking {
    return EntStatusTrackingFromJSONTyped(json, false);
}

export function EntStatusTrackingFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatusTracking {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntStatusTrackingEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'status': !exists(json, 'status') ? undefined : json['status'],
    };
}

export function EntStatusTrackingToJSON(value?: EntStatusTracking | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntStatusTrackingEdgesToJSON(value.edges),
        'id': value.id,
        'status': value.status,
    };
}


