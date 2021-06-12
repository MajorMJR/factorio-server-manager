import React, {useEffect, useState} from "react";
import savesResource from "../../../api/resources/saves";
import Panel from "../../components/Panel";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faDownload, faTrashAlt} from "@fortawesome/free-solid-svg-icons";
import Button from "../../components/Button";
import UploadSaveModal from "./components/UploadSaveModal";

const Saves = () => {

    const [saves, setSaves] = useState([]);
    const [isSaveUploadModalOpen, setIsSaveUploadModalOpen] = useState(false);

    const updateList = () => {
        savesResource.list()
            .then(setSaves)
    }

    useEffect(updateList, []);

    const deleteSave = save => {
        savesResource.delete(save)
            .then(updateList);
    }

    return (
        <Panel
            className="mb-4"
            title="Saves"
            actions={
                <>
                    <Button size="sm" onClick={() => setIsSaveUploadModalOpen(true)}>Upload File</Button>
                    <UploadSaveModal onSuccess={updateList} isOpen={isSaveUploadModalOpen} close={() => setIsSaveUploadModalOpen(false)}/>
                </>
            }
            content={
                <div className="overflow-x-auto w-full">
                    <table className="w-full">
                        <thead>
                        <tr className="text-left py-1">
                            <th>Name</th>
                            <th>Last Modified At</th>
                            <th>Size</th>
                            <th>Actions</th>
                        </tr>
                        </thead>
                        <tbody>
                        {saves.map(save =>
                            <tr className="py-2 md:py-1" key={save.name}>
                                <td className="pr-4">{save.name}</td>
                                <td className="pr-4">{(new Date(save.last_mod)).toISOString().replace('T', ' ').split('.')[0]}</td>
                                <td className="pr-4">{parseFloat(save.size / 1024 / 1024).toFixed(3)} MB</td>
                                <td>
                                    {save.name !== 'Load Latest' && <>
                                        <a href={`/api/saves/dl/${save.name}`} className="mr-2">
                                            <FontAwesomeIcon
                                                className="text-gray-light cursor-pointer hover:text-orange"
                                                icon={faDownload}/>
                                        </a>
                                        <FontAwesomeIcon className="text-red cursor-pointer hover:text-red-light mr-2"
                                                         onClick={() => deleteSave(save)} icon={faTrashAlt}/>
                                    </>}
                                </td>
                            </tr>
                        )}
                        </tbody>
                    </table>
                </div>
            }
        />
    )
}

export default Saves;