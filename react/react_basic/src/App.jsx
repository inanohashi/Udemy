import React, { useState } from "react";
import ColorfulMessage from './components/ColorfulMessage';

const App = () => {
    const [num, setNum] = useState(0);
    const [faceShowFlag, setFaceShowFlag] = useState(true);

    const onClickCountUp = () => {
        setNum(num + 1)
    }
    const onClickSwitchShowFlag = () => {
        setFaceShowFlag(!faceShowFlag)
    }

    return (
        <>
            <h1 style={{ color: 'red' }}>こんにちは!</h1>
            <ColorfulMessage color='blue'>お元気ですか？</ColorfulMessage>
            <ColorfulMessage color='pink'>元気です!</ColorfulMessage>
            <br />
            <button onClick={onClickSwitchShowFlag}>on/off</button>
            <button onClick={onClickCountUp}>カウントアップ</button>
            <p>{num}</p>
            {faceShowFlag && <p>(^o^)</p>}
        </>
    );
}

export default App;