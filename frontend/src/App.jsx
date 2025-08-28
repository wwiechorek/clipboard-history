import {useEffect, useRef, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { GetLatestClips, GetClipsAfter, HideApplication } from "../wailsjs/go/main/App";

function App() {
    const [clips, setClips] = useState([]);
    const lastTsRef = useRef('');

    // Buscar a hora atual quando o componente carregar
    useEffect(() => {
        GetLatestClips(0, 10).then(result => {
            setClips(result);
            if (result.length > 0) {
                lastTsRef.current = result[0].TSISO;
            }
        });

        const clipInterval = setInterval(() => {
            console.log('Checking for new clips after', lastTsRef.current);
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

    return (
        <div id="App">
            <div className="time-section">
                
                {clips.length > 0 && (
                    <div className="clips-section">
                        <h2>Clips Recentes</h2>
                        <ul>
                            {clips.map((clip, index) => (
                                <li key={index}>
                                    <strong>{new Date(clip.TSISO).toLocaleString()}:</strong> {clip.Content}
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
