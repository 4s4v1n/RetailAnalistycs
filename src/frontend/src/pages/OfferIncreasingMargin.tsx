import { FC, useState } from "react";
import { observer } from "mobx-react-lite";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { notifyError, notifySucces } from "../components/Notify/Notify";
import { saveAs } from 'file-saver'
import HeaderForm from "../components/FunctionHeader/FunctionHeader";
import { FUNCTION_MARGIN } from "../utils/const";
import View from "../components/View/View";
import { OfferIncreasingMarginResponse } from "../models/response/OfferIncreasingMarginResponse";
import OfferIncreasingMarginService from "../services/OfferIncreasingMarginService";

const OfferIncreasingMargin: FC = () => {

    const [data, setData] = useState<OfferIncreasingMarginResponse[]>();

    const queryClient = useQueryClient()

    const name = FUNCTION_MARGIN

    const { mutate: mutationForm } = useMutation({
        mutationFn: OfferIncreasingMarginService.get,
        onSuccess: (data) => {
            setData(data)
            notifySucces(`Successful response`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't response value in ${name}`);
        },
    })

    const { mutate: mutateExport } = useMutation({
        mutationFn: OfferIncreasingMarginService.export,
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
            <HeaderForm mutationForm={mutationForm} mutationExport={mutateExport} typeFunction={FUNCTION_MARGIN} />
            <View data={data} />
        </div>
    )
}

export default observer(OfferIncreasingMargin);