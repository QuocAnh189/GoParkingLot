import { Tabs, TabsList, TabsTrigger } from "@components/ui/tabs";

interface IInoutProps {
  stateInOut: any;
  setStateInOut: any;
}
export default function Inout(props: IInoutProps) {
  const { stateInOut, setStateInOut } = props;
  return (
    <div className="flex items-center space-x-2">
      <Tabs
        className="w-full"
        value={stateInOut}
        onValueChange={(value: any) => {
          setStateInOut(value);
        }}
      >
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger className="bg-gray text-white" value="check-out">
            Ra
          </TabsTrigger>
          <TabsTrigger className="bg-gray text-white" value="check-in">
            Vào
          </TabsTrigger>
        </TabsList>
      </Tabs>
    </div>
  );
}
