import React, { useEffect, useState } from 'react'


export default function App() {
const [videos, setVideos] = useState([])


useEffect(() => {
if (window.Telegram?.WebApp) {
window.Telegram.WebApp.init()
}


fetch('/api/videos')
.then(r => r.json())
.then(setVideos)
.catch(console.error)
}, [])


return (
<div style={{ padding: 20, fontFamily: 'Inter, sans-serif' }}>
<h1>Тренинг программа</h1>
{videos.length === 0 && <p>Загрузка...</p>}
{videos.map(v => (
<div key={v.id} style={{ marginBottom: 24 }}>
<h3>{v.title}</h3>
<div style={{ position: 'relative', paddingBottom: '56.25%', height: 0 }}>
<iframe
title={v.title}
src={`https://www.youtube.com/embed/${v.youtubeId}`}
style={{ position: 'absolute', top: 0, left: 0, width: '100%', height: '100%' }}
frameBorder="0"
allowFullScreen
/>
</div>
</div>
))}
</div>
)
}