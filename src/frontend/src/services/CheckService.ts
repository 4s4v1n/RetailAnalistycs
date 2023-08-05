import $api from "../http";
import { AxiosResponse } from "axios";
import { CheckResponse } from "../models/response/CheckResponse";

const tableName = "checks"

export default class CheckService {
    static async get(): Promise<CheckResponse[]> {
        const response = $api.get<CheckResponse[]>(`/table/${tableName}`)
        return (await response).data;
    }

    static async post(data: CheckResponse): Promise<AxiosResponse> {
        data.transaction_id = Number(data.transaction_id);
        data.sku_id = Number(data.sku_id);
        data.sku_amount = Number(data.sku_amount);
        data.sku_sum = Number(data.sku_sum);
        data.sku_sum_paid = Number(data.sku_sum_paid);
        data.sku_discount = Number(data.sku_discount);
        const response = $api.post(`/table/${tableName}`, data)
        return response
    }

    static async patch(data: CheckResponse): Promise<AxiosResponse> {
        data.transaction_id = Number(data.transaction_id);
        data.sku_id = Number(data.sku_id);
        data.sku_amount = Number(data.sku_amount);
        data.sku_sum = Number(data.sku_sum);
        data.sku_sum_paid = Number(data.sku_sum_paid);
        data.sku_discount = Number(data.sku_discount);
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
            data[i].transaction_id = Number(data[i].transaction_id);
            data[i].sku_id = Number(data[i].sku_id);
            data[i].sku_amount = Number(data[i].sku_amount);
            data[i].sku_sum = Number(data[i].sku_sum);
            data[i].sku_sum_paid = Number(data[i].sku_sum_paid);
            data[i].sku_discount = Number(data[i].sku_discount);
        }
        const response = $api.post(`/import/${tableName}`, data)
        return response
    }
}
