import { Link } from "react-router-dom";

import { useListMachines } from "@/api/hooks/machines";
import { Button } from "@/components/button";
import FormachEmptyScreen from "@/pages/dashboard/formach/components/empty-screen";
import ListFormach from "@/pages/dashboard/formach/components/list-formachs";

export function ListFormachPage() {
  const listMachines = useListMachines();

  if (listMachines.data?.length === 0) {
    return <FormachEmptyScreen />;
  }

  return (
    <section>
      <div className="flex justify-between py-6">
        <h3 className="font-medium text-lg">Formachs</h3>
        <Link to="/dashboard/formach/create">
          <Button size="sm">Create Formach</Button>
        </Link>
      </div>

      <ListFormach machines={listMachines.data} />
    </section>
  );
}
