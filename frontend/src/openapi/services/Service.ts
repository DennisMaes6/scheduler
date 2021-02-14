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
     * Sets the model paramters in the backend.
     * @param requestBody The model parameters to be set.
     * @returns any Models parameters updated succesfully.
     * @throws ApiError
     */
    public static async postService(
        requestBody: ModelParameters,
    ): Promise<any> {
        const result = await __request({
            method: 'POST',
            path: `/set-model-params`,
            body: requestBody,
        });
        return result.body;
    }

}