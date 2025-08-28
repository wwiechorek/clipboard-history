import {useEffect, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { GetClipboarText, GetLatestClips } from "../wailsjs/go/main/App";

function App() {
    const [currentTime, setCurrentTime] = useState('');
    const [clips, setClips] = useState([]);
    
    const updateCurrentTime = (time) => setCurrentTime(time);

    function getTimeFromGo() {
        GetClipboarText().then(updateCurrentTime);
    }

    // Buscar a hora atual quando o componente carregar
    useEffect(() => {
        setInterval(() => {
            getTimeFromGo();
        }, 1000)

        setInterval(() => {
            GetLatestClips(0, 10).then(clips => {
                setClips(clips)
            });
        }, 1000)
    }, []);

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            
            <div className="time-section">
                <div className="time-display">{currentTime}</div>
                
                {clips.length > 0 && (
                    <div className="clips-section">
                        <h2>Clips Recentes</h2>
                        <ul>
                            {clips.map((clip, index) => (
                                <li key={index}>
                                    <strong>{new Date(clip.TsIso).toLocaleString()}:</strong> {clip.Content}
                                </li>
                            ))}
                        </ul>
                    </div>
            )}
            </div>
            
        </div>
    )
}

export default App
