import { Table } from 'flowbite-react';
import { FC, useContext, useState } from 'react';
import { Context } from '../..';
import UpdateModal from '../Modal/UpdateModel';
import DeleteModal from '../Modal/DeleteModal';

export interface StandardComponentProps {
  data?: any[]
  mutationUpdate: any
  mutationDelete: any
}

const TableCustom: FC<StandardComponentProps> = ({ data, mutationUpdate, mutationDelete }) => {
  const { userStore } = useContext(Context)
  const [openModalUpdate, setOpenModalUpdate] = useState<string | undefined>();
  const [openModalDelete, setOpenModalDelete] = useState<string | undefined>();
  const [selectedValueDelete, setSelectedValueDelete] = useState(null);
  const [selectedValueUpdate, setSelectedValueUpdate] = useState<[string, unknown][]>([]);

  const handleRemoveClick = (value: any) => {
    setSelectedValueDelete(value);
  }

  const handleUpdateClick = (value: any) => {
    if (data === undefined || data === null) {
      return
    }
    const dataItem = Object.entries(data[0]);
    dataItem.forEach((element, i) => {
      element[1] = value[i]
    });
    setSelectedValueUpdate(dataItem);
  }


  return (
    <div className="p-4 sm:ml-64">
      <DeleteModal
        openModalDelete={openModalDelete}
        setOpenModalDelete={setOpenModalDelete}
        mutationDelete={mutationDelete}
        selectedValue={selectedValueDelete}
      />
      <UpdateModal
        openModalUpdate={openModalUpdate}
        setOpenModalUpdate={setOpenModalUpdate}
        keys={selectedValueUpdate}
        mutationUpdate={mutationUpdate}
      />
      <Table>
        <Table.Head>
          {
            data !== undefined && data !== null ?
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
          {
            data !== undefined && data !== null ?
              data.map((item, i) => (
                <Table.Row key={i} className="bg-white dark:border-gray-700 dark:bg-gray-800">
                  {Object.values(item).map((value, j) => (
                    < Table.Cell key={j} className="whitespace-nowrap font-medium text-gray-900 dark:text-white" >
                      {String(value)}
                    </Table.Cell>
                  ))}
                  {userStore.role === "admin" ?
                    <td>
                      <Table.Cell>
                        <button
                          className="font-medium text-blue-600 hover:underline dark:text-cyan-500"
                          onClick={() => { setOpenModalUpdate("default"); handleUpdateClick(Object.values(item)) }}
                        >
                          <p>
                            Update
                          </p>
                        </button>
                      </Table.Cell>
                      <Table.Cell>
                        <button
                          className="font-medium text-red-600 hover:underline dark:text-cyan-500"
                          onClick={() => { setOpenModalDelete("default"); handleRemoveClick(Object.values(item)[0]) }}
                        >
                          Delete
                        </button>
                      </Table.Cell>
                    </td>
                    : null}
                </Table.Row>
              )) : null}
        </Table.Body>
      </Table>
    </div >
  )
}

export default TableCustom;