import React, {useState, useEffect} from "react";

const Checkbox = ({name, text, inputRef, checked, className, textSize = 'sm', onChange = null}) => {

    const [value, setValue] = useState(false);

    const updateChecked = event => {
        if(onChange) {
            event.persist();
            onChange(event);
        }
        setValue(event.target.checked)
    }

    useEffect(() => {
        if (typeof checked === 'boolean') {
            setValue(checked);
        }
    }, [checked]);

    return (
        <label className={`block text-gray-500 font-bold ${className}`}>
            <input
                className="mr-2 leading-tight"
                ref={inputRef}
                name={name}
                id={name}
                onChange={updateChecked}
                type="checkbox"
                checked={value}
            />
            <span className={`text-${textSize}`}>{text}</span>
        </label>
    )
}

export default Checkbox;