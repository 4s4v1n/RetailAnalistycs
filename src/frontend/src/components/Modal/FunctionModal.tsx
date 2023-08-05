import { Label, TextInput } from "flowbite-react";
import { FC, useState } from "react";
import { RxCross2 } from "react-icons/rx";


export interface ModalFunctionProps {
    openModalForm?: string | undefined
    setOpenModalForm: React.Dispatch<React.SetStateAction<string | undefined>>
    mutationForm: any;
    keys: [string, string][];
    setDataForm: React.Dispatch<React.SetStateAction<{[key: string]: string;}>>
}

const FunctionModal: FC<ModalFunctionProps> = ({ openModalForm, setOpenModalForm, keys, mutationForm, setDataForm }) => {
    const [formData, setFormData] = useState<{ [key: string]: string }>({});

    const acceptForm = () => {
        setDataForm(formData)
        mutationForm(formData)
    }

    return (
        <div>
            {
                openModalForm !== undefined ? (
                    <div className="flex justify-center items-center overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
                        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>
                        <div className="relative w-full max-w-md max-h-full">
                            <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
                                <button onClick={() => setOpenModalForm(undefined)}
                                    className="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white"
                                    data-modal-hide="authentication-modal">
                                    <RxCross2 className="w-5 h-5 fill-white" />
                                </button>
                                <div className="px-6 py-6 lg:px-8">
                                    <h3 className="mb-4 text-xl font-medium text-gray-900 dark:text-white">Add item</h3>
                                    <form className="space-y-6" action="#">
                                        {keys.map((key) => (
                                            <div key={key[0]}>
                                                <div className="mb-2 block">
                                                    <Label value={key[0]} />
                                                </div>
                                                <TextInput key={key[0]}
                                                    placeholder={key[1]}
                                                    onChange={(e) => setFormData({ ...formData, [key[0]]: e.target.value })}
                                                />
                                            </div>
                                        ))}
                                        <button
                                            className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                                            onClick={() => { acceptForm(); setOpenModalForm(undefined); }}>
                                            Accept
                                        </button>
                                        <button
                                            className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-red-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                                            onClick={() => setOpenModalForm(undefined)}>
                                            Decline
                                        </button>
                                    </form>
                                </div>
                            </div>
                        </div >
                    </div >
                ) : null
            }
        </div>
    )
}

export default FunctionModal;