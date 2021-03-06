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
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
    EntStatusTracking,
    EntStatusTrackingFromJSON,
    EntStatusTrackingFromJSONTyped,
    EntStatusTrackingToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTreatmentEdges
 */
export interface EntTreatmentEdges {
    /**
     * 
     * @type {EntUser}
     * @memberof EntTreatmentEdges
     */
    doctorID?: EntUser;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntTreatmentEdges
     */
    patientID?: EntPatient;
    /**
     * 
     * @type {EntStatusTracking}
     * @memberof EntTreatmentEdges
     */
    statusID?: EntStatusTracking;
}

export function EntTreatmentEdgesFromJSON(json: any): EntTreatmentEdges {
    return EntTreatmentEdgesFromJSONTyped(json, false);
}

export function EntTreatmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTreatmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctorID': !exists(json, 'DoctorID') ? undefined : EntUserFromJSON(json['DoctorID']),
        'patientID': !exists(json, 'PatientID') ? undefined : EntPatientFromJSON(json['PatientID']),
        'statusID': !exists(json, 'StatusID') ? undefined : EntStatusTrackingFromJSON(json['StatusID']),
    };
}

export function EntTreatmentEdgesToJSON(value?: EntTreatmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'doctorID': EntUserToJSON(value.doctorID),
        'patientID': EntPatientToJSON(value.patientID),
        'statusID': EntStatusTrackingToJSON(value.statusID),
    };
}


