import { ReactNode, useEffect, useState } from 'react'

type Props = {
  loading: boolean
  children?: ReactNode | ReactNode[]
}

const Splash = ({ loading, children }: Props) => {
  const maxValue = 100

  const [currentValue, setCurrentValue] = useState(0)

  useEffect(() => {
    const interval = setInterval(() => {
      if (currentValue < maxValue) {
        setCurrentValue((prev) => prev + 1)
      }
    }, 5)

    return () => clearInterval(interval)
  }, [currentValue])

  if (loading) {
    return (
      <div className=" flex size-full h-screen items-center justify-center bg-slate-800">
        <div className="flex w-full flex-col items-center gap-12">
          <img
            src="/images/terraspect-logo.svg"
            alt="Terraspect Logo"
            width={150}
            height={150}
          />
          <div
            className="flex h-4 w-1/4 overflow-hidden rounded-full bg-white"
            role="progressbar"
          >
            <div
              className="flex flex-col justify-center overflow-hidden whitespace-nowrap rounded-full bg-red-600 text-center text-xs text-white transition duration-500"
              style={{ width: `${currentValue}%` }}
            />
          </div>
        </div>
      </div>
    )
  }

  return <>{children}</>
}

export default Splash
