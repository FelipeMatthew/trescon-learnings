import type { ReactNode } from "react"

const Title = (children : ReactNode) => {
  return (
    <h3 className="font-bold text-blue-900 text-xl mb-2 uppercase">{children}</h3>
  )
}

export default Title