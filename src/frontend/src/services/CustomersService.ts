import $api from "../http";
import { AxiosResponse } from "axios";
import { CustomersResponse } from "../models/response/CustomersResponse";

const tableName = "customers"

export default class CustomersService {
    static async get(): Promise<CustomersResponse[]> {
        const response = $api.get<CustomersResponse[]>(`/view/${tableName}`)
        return (await response).data;
    }

    static async export(): Promise<AxiosResponse> {
        const response = $api.get(`/export/${tableName}`)
        return response
    }
}