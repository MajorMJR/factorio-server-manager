import React from "react";

const Panel = ({title, content, actions, className}) => {
    return (
        <div className={(className ? className :'') + ' accentuated rounded-sm bg-gray-dark shadow-xl pb-4'}>
            <div className="px-4 py-2 text-xl text-dirty-white font-bold">
                {title}
            </div>
            <div className="text-white rounded-sm bg-gray-medium shadow-inner mx-4 px-6 pt-4 pb-6">
                {content}
            </div>
            {actions &&
                <div className="mx-4 pt-4">
                    {actions}
                </div>
            }
        </div>
    )
}

export default Panel;