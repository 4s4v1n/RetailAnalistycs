import { FC } from "react";
import { observer } from "mobx-react-lite";
import TableCustom from "../components/Table/Table";
import Header from "../components/Header/Header";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { notifyError, notifySucces } from "../components/Notify/Notify";
import { saveAs } from 'file-saver'
import TransactionService from "../services/TransactionService";

const Transaction: FC = () => {

    const queryClient = useQueryClient()

    const name = "transactions"

    const { data } = useQuery([{ name }], () => TransactionService.get(),)

    const { mutate: mutationAdd } = useMutation({
        mutationFn: TransactionService.post,
        onSuccess: () => {
            notifySucces(`Successful add`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't add value in ${name}`);
        },
    })

    const { mutate: mutationUpdate } = useMutation({
        mutationFn: TransactionService.patch,
        onSuccess: () => {
            notifySucces(`Successful update`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't update value in ${name}`);
        },
    })

    const { mutate: mutationDelete } = useMutation({
        mutationFn: TransactionService.delete,
        onSuccess: () => {
            notifySucces(`Successful delete`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't delete value in ${name}`);
        },
    })

    const { mutate: mutateExport } = useMutation({
        mutationFn: TransactionService.export,
        onSuccess: (data) => {
            notifySucces(`Successful export`)
            const blob = new Blob([data.data], { type: 'text/csv;charset=utf-8;' });
            saveAs(blob, `${name}.csv`);
        },
        onError: () => {
            notifyError(`Can't export ${name}`);
        },
    })

    const { mutate: mutateImport } = useMutation({
        mutationFn: TransactionService.import,
        onSuccess: () => {
            notifySucces(`Successful import`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't import ${name}`);
        },
    })


    return (
        <div>
            <Header header={name} mutationImport={mutateImport} mutationAdd={mutationAdd} mutationExport={mutateExport} keys={data !== undefined ? data[0] !== undefined ? Object.keys(data[0]) : [] : []} />
            <TableCustom data={data} mutationUpdate={mutationUpdate} mutationDelete={mutationDelete} />
        </div>
    )
}

export default observer(Transaction);