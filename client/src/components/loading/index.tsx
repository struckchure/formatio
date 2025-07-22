import { Spinner } from "@/components/icons/spinner";

export function LoadingScreen() {
  return (
    <div className="w-full h-full flex-center">
      <div className="flex items-end">
        <h5 className="text-5xl font-medium animate-pulse">Formati</h5>
        <Spinner size={32} color="#ddd" />
      </div>
    </div>
  );
}
