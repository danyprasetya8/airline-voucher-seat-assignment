import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import DatePicker from '@/components/DatePicker'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import { useState } from 'react'
import { format } from 'date-fns'
import { checkVoucher, generateVoucher } from '@/api/voucher'
import { toast } from 'sonner'
import axios from 'axios'
import type { ErrorResponse } from '@/interface/response'

const AIRCRAFT_TYPES = ['ATR', 'Airbus 320', 'Boeing 737 Max']

function App() {
  const [date, setDate] = useState<Date | undefined>()
  const [aircraftType, setAircraftType] = useState<string | undefined>()
  const [crewId, setCrewId] = useState<string | undefined>()
  const [crewName, setCrewName] = useState<string | undefined>()
  const [flightNumber, setFlightNumber] = useState<string | undefined>()

  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    try {
      const { data: checkRes } = await checkVoucher({
        date: date ? format(date, 'yyyy-MM-dd') : '',
        flightNumber
      })

      if (checkRes?.data.exists) {
        toast.error(('voucher already exist'))
        return
      }

      const { data: generateRes } = await generateVoucher({
        id: crewId,
        name: crewName,
        flightNumber,
        date: date ? format(date, 'yyyy-MM-dd') : '',
        aircraft: aircraftType
      })

      if (!generateRes.data.success) {
        toast.error('failed to generate voucher')
        return
      }

      toast.success('Voucher created on seats: ' + generateRes.data.seats.join(', '))

    } catch(err) {
      if (axios.isAxiosError<ErrorResponse>(err)) {
        toast.error(err.response?.data.error)
      } else if (err instanceof Error) {
        toast.error(err.message)
      } else {
        toast.error('Something went wrong')
      }
    }
  }

  return (
    <div className="w-1/3 mx-auto my-30">
      <div className="text-3xl font-bold mb-8">
        Airline Voucher Seat Assignment
      </div>

      <form
        onSubmit={submit}
        className="grid grid-cols-2 gap-8"
      >
        <div>
          <div className="mb-2">
            Crew ID
          </div>
          <Input
            placeholder="Input crew id"
            onChange={e => setCrewId(e.target.value)}
          />
        </div>

        <div>
          <div className="mb-2">
            Crew Name
          </div>
          <Input
            placeholder="Input crew name"
            onChange={e => setCrewName(e.target.value)}
          />
        </div>

        <div>
          <div className="mb-2">
            Flight Number
          </div>
          <Input
            placeholder="Input flight number"
            onChange={e => setFlightNumber(e.target.value)}
          />
        </div>

        <div>
          <div className="mb-2">
            Flight Date
          </div>
          <DatePicker
            date={date}
            onSelectDate={d => setDate(d)}
          />
        </div>

        <div>
          <div className="mb-2">
            Aircraft Type
          </div>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="outline" className="w-full text-left justify-start font-normal">
                {aircraftType ? aircraftType : 'Choose'}
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="w-56" align="start">
              <DropdownMenuGroup>
                {AIRCRAFT_TYPES.map(at => (
                  <DropdownMenuItem key={at} onClick={() => setAircraftType(at)}>
                    {at}
                  </DropdownMenuItem>
                ))}
              </DropdownMenuGroup>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>

        <div />

        <Button
          className="col-span-2 cursor-pointer"
          type="submit"
        >
          Generate
        </Button>
      </form>
    </div>
  )
}

export default App
