/// <reference types="vite/client" />
/// <reference types="vite-plugin-svgr/client" />
declare module "*.png";
declare module "*.svg" {
  import * as React from "react";

  export const ReactComponent: React.FunctionComponent<
    React.SVGProps<SVGSVGElement> & { title?: string }
  >;

  export default ReactComponent;
}

declare module "*.jpeg";
declare module "*.jpg";
