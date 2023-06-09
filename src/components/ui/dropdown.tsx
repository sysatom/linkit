import React from "react";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { CheckIcon, CherryIcon } from "lucide-react"

interface IDropdown {
  triggerLabel?: string;
  isLoading?: boolean;
  items?: {
    label: string;
    id: string;
  }[];
  selectedId?: string;
  handleOnClick: (id: string) => void;
  scrollRef?: React.RefObject<HTMLDivElement>;
}

const Dropdown = ({
                    triggerLabel,
                    items,
                    selectedId,
                    handleOnClick,
                    scrollRef,
                    isLoading,
                  }: IDropdown) => {
  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger asChild>
        {isLoading ? (
          <CherryIcon className="animate-spin text-zinc-400 dark:text-zinc-500" />
        ) : (
          <button
            className="py-1 px-4 border max-w-[256px] truncate outline-none rounded border-zinc-300 dark:border-zinc-700 bg-zinc-100 dark:bg-zinc-800 font-sans text-sm"
            aria-label={triggerLabel}
          >
            {triggerLabel}
          </button>
        )}
      </DropdownMenu.Trigger>

      <DropdownMenu.Portal>
        <DropdownMenu.Content
          className="shadow-sm bg-zinc-100 max-h-[250px] overflow-y-auto min-w-[192px] max-w-[256px] dark:bg-zinc-800 border border-zinc-300 dark:border-zinc-700 rounded-md p-2"
          sideOffset={4}
          align="end"
        >
          {items?.map((item) => (
            <DropdownMenu.Item
              key={item.id}
              className={`px-2 flex items-center justify-between cursor-default py-1 outline-none rounded bg-transparent ring-0 hover:bg-zinc-200 hover:dark:bg-zinc-700 text-sm hover:text-zinc-900 dark:hover:text-zinc-50 ${
                selectedId === item.id
                  ? "text-zinc-900 dark:text-zinc-50"
                  : "text-zinc-500 dark:text-zinc-400"
              }`}
              onClick={() => handleOnClick(item.id)}
            >
              {item.label}{" "}
              {selectedId === item.id && (
                <CheckIcon className="text-zinc-900 dark:text-white" />
              )}
            </DropdownMenu.Item>
          ))}
          <div ref={scrollRef} />
        </DropdownMenu.Content>
      </DropdownMenu.Portal>
    </DropdownMenu.Root>
  );
};

export default Dropdown;
