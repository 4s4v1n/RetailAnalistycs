import $api from "../http";
import { PersonalInfoResponse } from "../models/response/PersonalInfoResponse";
import { AxiosResponse } from "axios";


const tableName = "personal_information"

export default class PersonalInfoService {
    static async get(): Promise<PersonalInfoResponse[]> {
        const response = $api.get<PersonalInfoResponse[]>(`/table/${tableName}`)
        return (await response).data;

    }

    static async post(data: PersonalInfoResponse): Promise<AxiosResponse> {
        const response = $api.post(`/table/${tableName}`, data)
        return response
    }

    static async patch(data: PersonalInfoResponse): Promise<AxiosResponse> {
        data.customer_id = Number(data.customer_id);
        const response = $api.patch(`/table/${tableName}`, data)
        return response
    }

    static async delete(data: string): Promise<AxiosResponse> {
        const response = $api.delete(`/table/${tableName}/${data}`)
        return response
    }

    static async export(): Promise<AxiosResponse> {
        const response = $api.get(`/export/${tableName}`)
        return response
    }

    static async import(data: any): Promise<AxiosResponse> {
        for (let i = 0; i < data.length; i++) {
            data[i].customer_id = Number(data[i].customer_id);
        }
        const response = $api.post(`/import/${tableName}`, data)
        return response
    }
}
