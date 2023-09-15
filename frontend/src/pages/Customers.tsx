import { FC } from "react";
import { observer } from "mobx-react-lite";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { notifyError, notifySucces } from "../components/Notify/Notify";
import { saveAs } from 'file-saver'
import CustomersService from "../services/CustomersService";
import View from "../components/View/View";
import HeaderView from "../components/View/HeaderVIew";

const Customers: FC = () => {

    const queryClient = useQueryClient()

    const name = "customers"

    const { data } = useQuery([{ name }], () => CustomersService.get(),)


    const { mutate: mutationUpdate } = useMutation({
        mutationFn: CustomersService.get,
        onSuccess: () => {
            notifySucces(`Successful update`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't update value in ${name}`);
        },
    })

    const { mutate: mutateExport } = useMutation({
        mutationFn: CustomersService.export,
        onSuccess: (data) => {
            notifySucces(`Successful export`)
            const blob = new Blob([data.data], { type: 'text/csv;charset=utf-8;' });
            saveAs(blob, `${name}.csv`);
        },
        onError: () => {
            notifyError(`Can't export ${name}`);
        },
    })


    return (
        <div>
            <HeaderView header={name} mutationExport={mutateExport} mutationUpdate={mutationUpdate}/>
            <View data={data} />
        </div>
    )
}

export default observer(Customers);