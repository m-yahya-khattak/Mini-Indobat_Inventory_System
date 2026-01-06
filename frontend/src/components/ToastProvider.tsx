'use client';

import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function ToastProvider({ children }: { children: React.ReactNode }) {
  return (
    <>
      {children}
      <ToastContainer
        position="top-right"
        autoClose={4000}
        hideProgressBar={false}
        newestOnTop={true}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="light"
        toastClassName="!bg-white !shadow-lg !rounded-lg !border !border-slate-200 !p-4 !min-h-[60px] !relative"
        bodyClassName="!m-0 !p-0 !pr-8 !flex !items-center !gap-3"
        progressClassName="!bg-gradient-to-r !from-blue-500 !to-indigo-500"
        closeButton={({ closeToast }) => (
          <button
            onClick={closeToast}
            className="!text-slate-400 hover:!text-slate-600 transition-colors !p-1 !-mr-2 !-mt-1"
            aria-label="Close"
          >
            <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clipRule="evenodd" />
            </svg>
          </button>
        )}
      />
    </>
  );
}

