import type React from "react";
import type { ReactNode } from "react"

interface Props {
  children: ReactNode;
}

const Title: React.FC<Props> = ({children}) => {
  return (
    <h3 className="font-bold text-blue-900 text-xl mb-2 uppercase">{children}</h3>
  )
}

export default Title