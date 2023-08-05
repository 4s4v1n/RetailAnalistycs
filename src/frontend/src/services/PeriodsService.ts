import $api from "../http";
import { AxiosResponse } from "axios";
import { PeriodsResponse } from "../models/response/PeriodsResponse";

const tableName = "periods"

export default class PeriodsService {
    static async get(): Promise<PeriodsResponse[]> {
        const response = $api.get<PeriodsResponse[]>(`/view/${tableName}`)
        return (await response).data;
    }

    static async export(): Promise<AxiosResponse> {
        const response = $api.get(`/export/${tableName}`)
        return response
    }
}