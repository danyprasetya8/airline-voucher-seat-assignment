import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'

function App() {

  return (
    <div className="w-1/3 mx-auto my-30">
      <div className="text-3xl font-bold mb-8">
        Airline Voucher Seat Assignment
      </div>

      <form
        action=""
        className="grid grid-cols-2 gap-8"
      >
        <div>
          <div className="mb-2">
            Crew ID
          </div>
          <Input placeholder="Input crew id" />
        </div>

        <div>
          <div className="mb-2">
            Crew Name
          </div>
          <Input placeholder="Input crew name" />
        </div>

        <div>
          <div className="mb-2">
            Flight Number
          </div>
          <Input placeholder="Input flight number" />
        </div>

        <div>
          <div className="mb-2">
            Flight Date
          </div>
          <Input placeholder="Input flight date" />
        </div>

        <div>
          <div className="mb-2">
            Aircraft Type
          </div>
          <Input placeholder="" />
        </div>

        <div />

        <Button
          className="col-span-2"
          type="submit"
        >
          Generate
        </Button>
      </form>
    </div>
  )
}

export default App
