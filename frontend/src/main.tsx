import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import Echo from './Pages/Echo.tsx'
import ChatBord from './Pages/ChatBord.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Echo />,
  },
  {
    path: "/chat",
    element: <ChatBord />,
  },
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
