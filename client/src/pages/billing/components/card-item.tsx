import _ from "lodash";

import { DbCardModel } from "@/api";

export function CardItem(props: DbCardModel) {
  return (
    <div
      className={`w-full p-2 ${
        props.isDefault === true ? "bg-secondary text-black" : "bg-primary"
      }`}
    >
      <label className="font-bold">
        {_.capitalize(props.cardType)} ending in {props.lastDigits}
      </label>
      <p>
        Expires {props.expiryMonth}/{props.expiryYear}
      </p>
    </div>
  );
}
