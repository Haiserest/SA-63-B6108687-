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


import * as runtime from '../runtime';
import {
    ControllersTreatment,
    ControllersTreatmentFromJSON,
    ControllersTreatmentToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientToJSON,
    EntStatusTracking,
    EntStatusTrackingFromJSON,
    EntStatusTrackingToJSON,
    EntTreatment,
    EntTreatmentFromJSON,
    EntTreatmentToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserToJSON,
} from '../models';

export interface CreatePatientRequest {
    patient: EntPatient;
}

export interface CreateStatustrackingRequest {
    statustracking: EntStatusTracking;
}

export interface CreateTreatmentRequest {
    treatment: ControllersTreatment;
}

export interface CreateUserRequest {
    user: EntUser;
}

export interface DeletePatientRequest {
    id: number;
}

export interface DeleteStatustrackingRequest {
    id: number;
}

export interface DeleteTreatmentRequest {
    id: number;
}

export interface DeleteUserRequest {
    id: number;
}

export interface GetPatientRequest {
    id: number;
}

export interface GetStatustrackingRequest {
    id: number;
}

export interface GetTreatmentRequest {
    id: number;
}

export interface GetUserRequest {
    id: number;
}

export interface ListPatientRequest {
    limit?: number;
    offset?: number;
}

export interface ListStatustrackingRequest {
    limit?: number;
    offset?: number;
}

export interface ListTreatmentRequest {
    limit?: number;
    offset?: number;
}

export interface ListUserRequest {
    limit?: number;
    offset?: number;
}

export interface UpdatePatientRequest {
    id: number;
    patient: EntPatient;
}

export interface UpdateStatustrackingRequest {
    id: number;
    statustracking: EntStatusTracking;
}

export interface UpdateTreatmentRequest {
    id: number;
    treatment: EntTreatment;
}

export interface UpdateUserRequest {
    id: number;
    user: EntUser;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create patient
     * Create patient
     */
    async createPatientRaw(requestParameters: CreatePatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.patient === null || requestParameters.patient === undefined) {
            throw new runtime.RequiredError('patient','Required parameter requestParameters.patient was null or undefined when calling createPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/patients`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntPatientToJSON(requestParameters.patient),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * Create patient
     * Create patient
     */
    async createPatient(requestParameters: CreatePatientRequest): Promise<EntPatient> {
        const response = await this.createPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create statustracking
     * Create statustracking
     */
    async createStatustrackingRaw(requestParameters: CreateStatustrackingRequest): Promise<runtime.ApiResponse<EntStatusTracking>> {
        if (requestParameters.statustracking === null || requestParameters.statustracking === undefined) {
            throw new runtime.RequiredError('statustracking','Required parameter requestParameters.statustracking was null or undefined when calling createStatustracking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/statustrackings`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntStatusTrackingToJSON(requestParameters.statustracking),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntStatusTrackingFromJSON(jsonValue));
    }

    /**
     * Create statustracking
     * Create statustracking
     */
    async createStatustracking(requestParameters: CreateStatustrackingRequest): Promise<EntStatusTracking> {
        const response = await this.createStatustrackingRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create treatment
     * Create treatment
     */
    async createTreatmentRaw(requestParameters: CreateTreatmentRequest): Promise<runtime.ApiResponse<ControllersTreatment>> {
        if (requestParameters.treatment === null || requestParameters.treatment === undefined) {
            throw new runtime.RequiredError('treatment','Required parameter requestParameters.treatment was null or undefined when calling createTreatment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/treatments`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ControllersTreatmentToJSON(requestParameters.treatment),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ControllersTreatmentFromJSON(jsonValue));
    }

    /**
     * Create treatment
     * Create treatment
     */
    async createTreatment(requestParameters: CreateTreatmentRequest): Promise<ControllersTreatment> {
        const response = await this.createTreatmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create user
     * Create user
     */
    async createUserRaw(requestParameters: CreateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling createUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * Create user
     * Create user
     */
    async createUser(requestParameters: CreateUserRequest): Promise<EntUser> {
        const response = await this.createUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get patient by ID
     * Delete a patient entity by ID
     */
    async deletePatientRaw(requestParameters: DeletePatientRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deletePatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get patient by ID
     * Delete a patient entity by ID
     */
    async deletePatient(requestParameters: DeletePatientRequest): Promise<object> {
        const response = await this.deletePatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * get statustracking by ID
     * Delete a statustracking entity by ID
     */
    async deleteStatustrackingRaw(requestParameters: DeleteStatustrackingRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteStatustracking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/statustrackings/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get statustracking by ID
     * Delete a statustracking entity by ID
     */
    async deleteStatustracking(requestParameters: DeleteStatustrackingRequest): Promise<object> {
        const response = await this.deleteStatustrackingRaw(requestParameters);
        return await response.value();
    }

    /**
     * get treatment by ID
     * Delete a treatment entity by ID
     */
    async deleteTreatmentRaw(requestParameters: DeleteTreatmentRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteTreatment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/treatments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get treatment by ID
     * Delete a treatment entity by ID
     */
    async deleteTreatment(requestParameters: DeleteTreatmentRequest): Promise<object> {
        const response = await this.deleteTreatmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUser(requestParameters: DeleteUserRequest): Promise<object> {
        const response = await this.deleteUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get patient by ID
     * Get a patient entity by ID
     */
    async getPatientRaw(requestParameters: GetPatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * get patient by ID
     * Get a patient entity by ID
     */
    async getPatient(requestParameters: GetPatientRequest): Promise<EntPatient> {
        const response = await this.getPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * get statustracking by ID
     * Get a statustracking entity by ID
     */
    async getStatustrackingRaw(requestParameters: GetStatustrackingRequest): Promise<runtime.ApiResponse<EntStatusTracking>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getStatustracking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/statustrackings/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntStatusTrackingFromJSON(jsonValue));
    }

    /**
     * get statustracking by ID
     * Get a statustracking entity by ID
     */
    async getStatustracking(requestParameters: GetStatustrackingRequest): Promise<EntStatusTracking> {
        const response = await this.getStatustrackingRaw(requestParameters);
        return await response.value();
    }

    /**
     * get treatment by ID
     * Get a treatment entity by ID
     */
    async getTreatmentRaw(requestParameters: GetTreatmentRequest): Promise<runtime.ApiResponse<EntTreatment>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getTreatment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/treatments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntTreatmentFromJSON(jsonValue));
    }

    /**
     * get treatment by ID
     * Get a treatment entity by ID
     */
    async getTreatment(requestParameters: GetTreatmentRequest): Promise<EntTreatment> {
        const response = await this.getTreatmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUserRaw(requestParameters: GetUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUser(requestParameters: GetUserRequest): Promise<EntUser> {
        const response = await this.getUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * list patient entities
     * List patient entities
     */
    async listPatientRaw(requestParameters: ListPatientRequest): Promise<runtime.ApiResponse<Array<EntPatient>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntPatientFromJSON));
    }

    /**
     * list patient entities
     * List patient entities
     */
    async listPatient(requestParameters: ListPatientRequest): Promise<Array<EntPatient>> {
        const response = await this.listPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * list statustracking entities
     * List statustracking entities
     */
    async listStatustrackingRaw(requestParameters: ListStatustrackingRequest): Promise<runtime.ApiResponse<Array<EntStatusTracking>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/statustrackings`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntStatusTrackingFromJSON));
    }

    /**
     * list statustracking entities
     * List statustracking entities
     */
    async listStatustracking(requestParameters: ListStatustrackingRequest): Promise<Array<EntStatusTracking>> {
        const response = await this.listStatustrackingRaw(requestParameters);
        return await response.value();
    }

    /**
     * list treatment entities
     * List treatment entities
     */
    async listTreatmentRaw(requestParameters: ListTreatmentRequest): Promise<runtime.ApiResponse<Array<EntTreatment>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/treatments`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntTreatmentFromJSON));
    }

    /**
     * list treatment entities
     * List treatment entities
     */
    async listTreatment(requestParameters: ListTreatmentRequest): Promise<Array<EntTreatment>> {
        const response = await this.listTreatmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * list user entities
     * List user entities
     */
    async listUserRaw(requestParameters: ListUserRequest): Promise<runtime.ApiResponse<Array<EntUser>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntUserFromJSON));
    }

    /**
     * list user entities
     * List user entities
     */
    async listUser(requestParameters: ListUserRequest): Promise<Array<EntUser>> {
        const response = await this.listUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * update patient by ID
     * Update a patient entity by ID
     */
    async updatePatientRaw(requestParameters: UpdatePatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updatePatient.');
        }

        if (requestParameters.patient === null || requestParameters.patient === undefined) {
            throw new runtime.RequiredError('patient','Required parameter requestParameters.patient was null or undefined when calling updatePatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntPatientToJSON(requestParameters.patient),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * update patient by ID
     * Update a patient entity by ID
     */
    async updatePatient(requestParameters: UpdatePatientRequest): Promise<EntPatient> {
        const response = await this.updatePatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * update statustracking by ID
     * Update a statustracking entity by ID
     */
    async updateStatustrackingRaw(requestParameters: UpdateStatustrackingRequest): Promise<runtime.ApiResponse<EntStatusTracking>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateStatustracking.');
        }

        if (requestParameters.statustracking === null || requestParameters.statustracking === undefined) {
            throw new runtime.RequiredError('statustracking','Required parameter requestParameters.statustracking was null or undefined when calling updateStatustracking.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/statustrackings/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntStatusTrackingToJSON(requestParameters.statustracking),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntStatusTrackingFromJSON(jsonValue));
    }

    /**
     * update statustracking by ID
     * Update a statustracking entity by ID
     */
    async updateStatustracking(requestParameters: UpdateStatustrackingRequest): Promise<EntStatusTracking> {
        const response = await this.updateStatustrackingRaw(requestParameters);
        return await response.value();
    }

    /**
     * update treatment by ID
     * Update a treatment entity by ID
     */
    async updateTreatmentRaw(requestParameters: UpdateTreatmentRequest): Promise<runtime.ApiResponse<EntTreatment>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateTreatment.');
        }

        if (requestParameters.treatment === null || requestParameters.treatment === undefined) {
            throw new runtime.RequiredError('treatment','Required parameter requestParameters.treatment was null or undefined when calling updateTreatment.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/treatments/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntTreatmentToJSON(requestParameters.treatment),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntTreatmentFromJSON(jsonValue));
    }

    /**
     * update treatment by ID
     * Update a treatment entity by ID
     */
    async updateTreatment(requestParameters: UpdateTreatmentRequest): Promise<EntTreatment> {
        const response = await this.updateTreatmentRaw(requestParameters);
        return await response.value();
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUserRaw(requestParameters: UpdateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateUser.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling updateUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUser(requestParameters: UpdateUserRequest): Promise<EntUser> {
        const response = await this.updateUserRaw(requestParameters);
        return await response.value();
    }

}
