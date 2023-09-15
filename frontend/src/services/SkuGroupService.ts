import $api from "../http";
import { AxiosResponse } from "axios";
import { SkuGroupResponse } from "../models/response/SkuGroupResponse";


const tableName = "sku_group"

export default class SkuGroupService {
    static async get(): Promise<SkuGroupResponse[]> {
        const response = $api.get<SkuGroupResponse[]>(`/table/${tableName}`)
        return (await response).data;
    }


    static async post(data: SkuGroupResponse): Promise<AxiosResponse> {
        data.group_id = Number(data.group_id);
        const response = $api.post(`/table/${tableName}`, data)
        return response
    }

    static async patch(data: SkuGroupResponse): Promise<AxiosResponse> {
        data.group_id = Number(data.group_id);
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
            data[i].group_id = Number(data[i].group_id);
        }
        const response = $api.post(`/import/${tableName}`, data)
        return response
    }
}
