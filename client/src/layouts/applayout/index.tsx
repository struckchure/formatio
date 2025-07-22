import { ChildrenProps } from "@/types";

export function AppLayout(props: ChildrenProps) {
  return (
    <div className="bg-primary text-white absolute top-0 left-0 w-full h-full overflow-auto">
      {props.children}
    </div>
  );
}
