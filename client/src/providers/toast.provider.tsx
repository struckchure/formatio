import { Toaster } from "react-hot-toast";
import { FaCheckCircle, FaExclamationCircle, FaSpinner } from "react-icons/fa";

export function ToastProvider() {
  return (
    <Toaster
      toastOptions={{
        error: {
          icon: <FaExclamationCircle size={15} />,
        },
        success: {
          icon: <FaCheckCircle size={15} />,
        },
        loading: {
          icon: (
            <span className="animate-spin">
              <FaSpinner size={10} />
            </span>
          ),
        },
        position: "top-center",
        className: "!rounded-xl !bg-white/5 !text-white",
      }}
    />
  );
}
