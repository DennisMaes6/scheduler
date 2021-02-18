/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ModelParameters } from '../models/ModelParameters';
import type { Schedule } from '../models/Schedule';
import { request as __request } from '../core/request';

export class Service {

    /**
     * Returns a generated schedule.
     * @returns Schedule A schedule.
     * @throws ApiError
     */
    public static async getService(): Promise<Schedule> {
        const result = await __request({
            method: 'GET',
            path: `/schedule`,
        });
        return result.body;
    }

    /**
     * Returns the current model parameters.
     * @returns ModelParameters The model parameters.
     * @throws ApiError
     */
    public static async getService1(): Promise<ModelParameters> {
        const result = await __request({
            method: 'GET',
            path: `/model-parameters/get`,
        });
        return result.body;
    }

    /**
     * Sets the model paramters in the backend.
     * @param requestBody The model parameters to be set.
     * @throws ApiError
     */
    public static async postService(
        requestBody: ModelParameters,
    ): Promise<void> {
        const result = await __request({
            method: 'POST',
            path: `/model-parameters/set`,
            body: requestBody,
        });
        return result.body;
    }

}