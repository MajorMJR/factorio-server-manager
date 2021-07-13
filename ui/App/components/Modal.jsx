import Panel from "./Panel";
import React from "react";
import * as ReactDom from "react-dom";

const modalRoot = document.getElementById('modal-root');

const Modal = ({title, content, isOpen, actions = null}) => {

    return ReactDom.createPortal((isOpen &&
        <div className="relative z-40">
            <div className="bg-black bg-opacity-75 fixed top-0 left-0 z-10 w-full min-h-screen">
                <Panel
                    title={title}
                    className="w-1/3 mx-auto mt-6"
                    content={content}
                    actions={actions}
                />
            </div>
        </div>), modalRoot)
}

export default Modal;