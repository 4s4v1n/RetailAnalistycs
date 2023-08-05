import { FC } from "react";
import { observer } from "mobx-react-lite";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { notifyError, notifySucces } from "../components/Notify/Notify";
import { saveAs } from 'file-saver'
import View from "../components/View/View";
import HeaderView from "../components/View/HeaderVIew";
import PeriodsService from "../services/PeriodsService";

const Periods: FC = () => {

    const queryClient = useQueryClient()

    const name = "periods"

    const { data } = useQuery([{ name }], () => PeriodsService.get(),)


    const { mutate: mutationUpdate } = useMutation({
        mutationFn: PeriodsService.get,
        onSuccess: () => {
            notifySucces(`Successful update`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't update value in ${name}`);
        },
    })

    const { mutate: mutateExport } = useMutation({
        mutationFn: PeriodsService.export,
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

export default observer(Periods);