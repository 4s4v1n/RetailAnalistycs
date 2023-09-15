import { Table } from 'flowbite-react';
import { FC, useState } from 'react';
import { DateAnalysingResponse } from '../../models/response/DateAnalysingResponse';
import UpdateDate from '../Modal/UpdateDate';

export interface StandardComponentProps {
    data?: DateAnalysingResponse[]
    mutationUpdate: any
}

const TableOnce: FC<StandardComponentProps> = ({ data, mutationUpdate }) => {
    const [openModalUpdate, setOpenModalUpdate] = useState<string | undefined>();

    return (
        <div className="p-4 sm:ml-64">
            <UpdateDate
                openModalUpdate={openModalUpdate}
                setOpenModalUpdate={setOpenModalUpdate}
                keys={data?.at(0)}
                mutationUpdate={mutationUpdate}
            />
            <Table>
                <Table.Head>
                    <Table.HeadCell>
                        Date
                    </Table.HeadCell>
                    <Table.HeadCell>
                        <span className="sr-only"></span>
                    </Table.HeadCell>
                </Table.Head>
                <Table.Body className="divide-y">
                    <Table.Row className="bg-white dark:border-gray-700 dark:bg-gray-800">
                        < Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white" >
                            {data?.at(0)?.date}
                        </Table.Cell>
                        <Table.Cell>
                            <button
                                className="font-medium text-blue-600 hover:underline dark:text-cyan-500"
                                onClick={() => setOpenModalUpdate("default")}>
                                Update
                            </button>
                        </Table.Cell>
                    </Table.Row>
                </Table.Body>
            </Table>
        </div >
    )
}

export default TableOnce;