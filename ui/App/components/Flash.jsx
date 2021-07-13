import React, { useEffect, useState } from 'react';
import Bus from '../../notifications';

export const Flash = () => {
    let [visibility, setVisibility] = useState(false);
    let [message, setMessage] = useState('');
    let [color, setColor] = useState('');

    let flashListener = ({message, color}) => {
        setVisibility(true);
        setMessage(message);
        setColor(color);
        setTimeout(() => {
            setVisibility(false);
        }, 4000);
    }

    useEffect(() => {
        Bus.addListener('flash', flashListener);

        return function () {
            Bus.removeListener('flash', flashListener);
        }
    }, []);

    return (
        visibility && <div onClick={() => setVisibility(false)} className={`bg-${color} cursor-pointer accentuated rounded fixed bottom-0 right-0 mr-8 mb-8 px-4 py-2 z-50 max-w-1/2 lg:max-w-none`}>
            <p>{message}</p>
        </div>
    )
}