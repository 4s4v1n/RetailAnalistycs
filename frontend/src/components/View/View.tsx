import { Table } from 'flowbite-react';
import { FC, useContext } from 'react';
import { Context } from '../..';

export interface StandardComponentProps {
    data?: any[]
}

const View: FC<StandardComponentProps> = ({ data }) => {
    const { userStore } = useContext(Context)

    return (
        <div className="p-4 sm:ml-64">
            <Table>
                <Table.Head>
                    {data !== undefined && data !== null ?
                        data[0] !== undefined ?
                            Object.keys(data[0]).map((item, i) => (
                                <Table.HeadCell key={i}>
                                    {item}
                                </Table.HeadCell>
                            )) : null : null
                    }
                    {userStore.role === "admin" ?
                        <Table.HeadCell>
                            <span className="sr-only"></span>
                        </Table.HeadCell>
                        : null}
                </Table.Head>
                <Table.Body className="divide-y">
                    {data !== undefined && data !== null?
                        data.map((item, i) => (
                            <Table.Row key={i} className="bg-white dark:border-gray-700 dark:bg-gray-800">
                                {Object.values(item).map((value, j) => (
                                    < Table.Cell key={j} className="whitespace-nowrap font-medium text-gray-900 dark:text-white" >
                                        {String(value)}
                                    </Table.Cell>
                                ))}
                            </Table.Row>
                        )) : null}
                </Table.Body>
            </Table>
        </div >
    )
}

export default View;