import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router } from "react-router-dom"
import { HelmetProvider } from 'react-helmet-async';
import { Analytics } from "@vercel/analytics/react"
import './index.css'
import App from './App.jsx'

createRoot(document.getElementById('root')).render(
  <HelmetProvider>
    <Router>
      <App />
      <Analytics />
    </Router>
  </HelmetProvider>,
)
