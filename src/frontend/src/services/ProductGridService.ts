import $api from "../http";
import { AxiosResponse } from "axios";
import { ProductGridResponse } from "../models/response/ProductGridResponse";

const tableName = "product_grid"

export default class ProductGridService {
    static async get(): Promise<ProductGridResponse[]> {
        const response = $api.get<ProductGridResponse[]>(`/table/${tableName}`)
        return (await response).data;
    }

    static async post(data: ProductGridResponse): Promise<AxiosResponse> {
        data.group_id = Number(data.group_id);
        data.sku_id = Number(data.sku_id);
        const response = $api.post(`/table/${tableName}`, data)
        return response
    }

    static async patch(data: ProductGridResponse): Promise<AxiosResponse> {
        data.group_id = Number(data.group_id);
        data.sku_id = Number(data.sku_id);
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
            data[i].sku_id = Number(data[i].sku_id);
        }
        const response = $api.post(`/import/${tableName}`, data)
        return response
    }
}
