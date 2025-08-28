import {useEffect, useRef, useState} from 'react';
import style from './App.module.css';
import { GetLatestClips, GetClipsAfter, HideApplication, PreventCopyText } from "../wailsjs/go/main/App";

function App() {
    const [clips, setClips] = useState([]);
    const lastTsRef = useRef('');

    useEffect(() => {
        GetLatestClips(0, 10).then(result => {
            setClips(result);
            if (result.length > 0) {
                lastTsRef.current = result[0].TSISO;
            }
        });

        const clipInterval = setInterval(() => {
            if (!lastTsRef.current) {
                return;
            }
            GetClipsAfter(lastTsRef.current).then(newClips => {
                if (newClips.length > 0) {
                    setClips(prev => [...newClips, ...prev]);
                    lastTsRef.current = newClips[0].TSISO;
                }
            });
        }, 1000);

        return () => {
            clearInterval(clipInterval);
        };
    }, []);

    useEffect(() => {
        window.addEventListener('blur', HideApplication);
        return () => {
            window.removeEventListener('blur', HideApplication);
        }
    }, [])

    const copyText = (text) => () => {
        PreventCopyText(text);
        navigator.clipboard.writeText(text).then(() => {
            HideApplication();
        });
    }

    return (
        <div id="app">
            <div className={style["time-section"]}>
                {clips.length > 0 && (
                    <div className={style["clips-section"]}>
                        <ul>
                            {clips.map((clip, index) => (
                                <li key={index} onClick={copyText(clip.Content)}>
                                    <div className={style["content"]}>{clip.Content}</div>
                                    <strong className={style["date"]}>{new Date(clip.TSISO).toLocaleString()}</strong>
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
