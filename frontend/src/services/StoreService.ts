import $api from "../http";
import { AxiosResponse } from "axios";
import { StoreResponse } from "../models/response/StoreResponse";


const tableName = "stores"

export default class StoreService {
    static async get(): Promise<StoreResponse[]> {
        const response = $api.get<StoreResponse[]>(`/table/${tableName}`)
        return (await response).data;
    }

    static async post(data: StoreResponse): Promise<AxiosResponse> {
        data.transaction_store_id = Number(data.transaction_store_id);
        data.sku_id = Number(data.sku_id);
        data.sku_purchase_price = Number(data.sku_purchase_price);
        data.sku_retail_price = Number(data.sku_retail_price);
        const response = $api.post(`/table/${tableName}`, data)
        return response
    }

    static async patch(data: StoreResponse): Promise<AxiosResponse> {
        data.transaction_store_id = Number(data.transaction_store_id);
        data.sku_id = Number(data.sku_id);
        data.sku_purchase_price = Number(data.sku_purchase_price);
        data.sku_retail_price = Number(data.sku_retail_price);
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
            data[i].transaction_store_id = Number(data[i].transaction_store_id);
            data[i].sku_id = Number(data[i].sku_id);
            data[i].sku_purchase_price = Number(data[i].sku_purchase_price);
            data[i].sku_retail_price = Number(data[i].sku_retail_price);
        }
        const response = $api.post(`/import/${tableName}`, data)
        return response
    }
}
