/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { InstanceData } from '../models/InstanceData';
import type { ModelParameters } from '../models/ModelParameters';
import type { Schedule } from '../models/Schedule';
import { request as __request } from '../core/request';

export class Service {

    /**
     * Returns a schedule generated with MiniZinc.
     * @returns Schedule A schedule.
     * @throws ApiError
     */
    public static async getSchedule(): Promise<Schedule> {
        const result = await __request({
            method: 'GET',
            path: `/schedule`,
        });
        return result.body;
    }

    /**
     * Returns the schedule as found in the db.
     * @returns Schedule A schedule.
     * @throws ApiError
     */
    public static async getDbSchedule(): Promise<Schedule> {
        const result = await __request({
            method: 'GET',
            path: `/db-schedule`,
        });
        return result.body;
    }

      /**
     * Returns the schedule as found in a file.
     * @returns Schedule A schedule.
     * @throws ApiError
     */
    public static async getFileSchedule(): Promise<Schedule> {
        const result = await __request({
            method: 'GET',
            path: `/file-schedule`,
        });
        return result.body;
    }

    /**
     * Returns the current model parameters.
     * @returns ModelParameters Succesful operation.
     * @throws ApiError
     */
    public static async getModelParams(): Promise<ModelParameters> {
        const result = await __request({
            method: 'GET',
            path: `/model-parameters/get`,
        });
        return result.body;
    }

    /**
     * Sets the model paramters in the backend.
     * @param requestBody The model parameters to be set.
     * @returns any Model parameters succesfully updated in backend.
     * @throws ApiError
     */
    public static async postModelParams(
        requestBody: ModelParameters,
    ): Promise<any> {
        const result = await __request({
            method: 'POST',
            path: `/model-parameters/set`,
            body: requestBody,
        });
        return result.body;
    }

    /**
     * Returns the current instance data.
     * @returns InstanceData Succesful operation.
     * @throws ApiError
     */
    public static async getInstanceData(): Promise<InstanceData> {
        const result = await __request({
            method: 'GET',
            path: `/instance-data/get`,
        });
        return result.body;
    }

    /**
     * Sets the insatnce data in the backend.
     * @param requestBody The instance data to be set.
     * @returns any Instance data succesfully updated in backend.
     * @throws ApiError
     */
    public static async postInstanceData(
        requestBody: InstanceData,
    ): Promise<any> {
        const result = await __request({
            method: 'POST',
            path: `/instance-data/set`,
            body: requestBody,
        });
        return result.body;
    }

}