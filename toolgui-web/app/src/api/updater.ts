var sessionID = ''

var updateSock: WebSocket | null = null
var healthSock: WebSocket | null = null

function getSocketURI() {
  var scheme = 'ws'
  if (window.location.origin.startsWith('https')) {
    scheme = 'wss'
  }

  return `${scheme}://${window.location.host}`
}

function getUpdateURI(pageName: string) {
  return `${getSocketURI()}/api/update/${pageName}`
}

function getHealthURI(pageName: string) {
  return `${getSocketURI()}/api/health/${pageName}`
}

export function wsUpdate(
  pageName: string,
  event: any,
  clearContainer: () => void,
  clearSession: () => void,
  recvNotifyPack: (pack: any) => void,
  finishUpdate: (pack: any) => void) {

  if (sessionID !== '') {
    event['session_id'] = sessionID
  }

  if (updateSock) {
    updateSock.close()
  }

  updateSock = new WebSocket(getUpdateURI(pageName))
  var jsonEvent = JSON.stringify(event)

  updateSock.onopen = function () {
    clearContainer()
    if (updateSock) {
      updateSock.send(jsonEvent)
    }
  }

  updateSock.onmessage = function (e) {
    const data = JSON.parse(e.data)
    if (data.session_id) {
      sessionID = data.session_id
      clearSession()
      return
    }

    if (data.success !== undefined) {
      finishUpdate(data)
      return
    }

    recvNotifyPack(data)
  }
}

export function initHealthSock(pageName: string) {
  if (healthSock) {
    return
  }

  healthSock = new WebSocket(getHealthURI(pageName))
  healthSock.onopen = function () {
    console.log('Start health beating')
  }

  // health beat / 1 mins
  setInterval(function () {
    if (healthSock) {
      healthSock.send(JSON.stringify({ session_id: sessionID }))
    }
  }, 60 * 1000);
}