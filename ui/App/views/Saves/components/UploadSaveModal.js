import Button from "../../../components/Button";
import React, {useState} from "react";
import {useForm} from "react-hook-form";
import saves from "../../../../api/resources/saves";
import Error from "../../../components/Error";
import Label from "../../../components/Label";
import Modal from "../../../components/Modal";

const defaultFileName = "Select File ..."

const UploadSaveModal = ({onSuccess, close, isOpen}) => {
    const {register, handleSubmit, errors} = useForm();
    const [fileName, setFileName] = useState(defaultFileName);
    const [isUploading, setIsUploading] = useState(false);

    const onSubmit = data => {
        setIsUploading(true);
        saves
            .upload(data.savefile[0])
            .then(message => {
                flash(message, "green");
                onSuccess();
                close();
            })
            .finally(() => {
                setFileName(defaultFileName);
                setIsUploading(false);
            });
    };

    return (
            <Modal
                title="Upload Save"
                isOpen={isOpen}
                onSubmit={handleSubmit(onSubmit)}
                content={
                    <div className="mb-6">
                        <Label text="Save File" htmlFor="savefile"/>
                        <div className="relative rounded-sm bg-gray-light shadow text-black w-full">
                            <input
                                className="absolute  left-0 top-0 opacity-0 cursor-pointer w-full h-full"
                                ref={register({required: true})}
                                onChange={e => setFileName(e.currentTarget.files[0].name)}
                                name="savefile"
                                id="savefile" type="file"/>
                            <div className="px-3 py-2">{fileName}</div>
                        </div>
                        <Error error={errors.savefile} message="Savefile is required"/>
                    </div>
                }
                actions={
                    <>
                        <Button size="sm" type="success"
                                isLoading={isUploading}
                                isDisabled={fileName === defaultFileName}
                                isSubmit>Upload</Button>
                        <Button size="sm" className="ml-1" onClick={close}>Cancel</Button>
                    </>
                }
            />
    )
}

export default UploadSaveModal;