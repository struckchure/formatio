import { useRouter } from "@/api/hooks/router";
import { Button } from "@/components/button";
import { ArrowLeftIcon } from "@/components/icons/arrow-left";

export function BackButton() {
  const router = useRouter();

  return (
    <Button
      theme="ghost"
      icon={<ArrowLeftIcon />}
      onClick={() => router.goBack()}
      className="px-0 text-white"
    >
      back
    </Button>
  );
}
