package businesssdk

// This class extends the CAPI event class and provides more functionality for setting hashed values

func (o *CapiEvent) SetEmail(unhashedEmail string) bool {
	if str, ok := NormAndHashStr(unhashedEmail); ok {
		o.SetHashedEmail(*str)
		return ok
	}
	return false
}

func (o *CapiEvent) SetMobileAdId(unhashedMobileAdId string) bool {
	if str, ok := NormAndHashStr(unhashedMobileAdId); ok {
		o.SetHashedMobileAdId(*str)
		return ok
	}
	return false
}

func (o *CapiEvent) SetIdfv(unhashedIdfv string) bool {
	if str, ok := NormAndHashStr(unhashedIdfv); ok {
		o.SetHashedIdfv(*str)
		return ok
	}
	return false
}

func (o *CapiEvent) SetPhoneNumber(unhashedPhoneNumber string) bool {
	if str, ok := NormAndHashPhoneNum(unhashedPhoneNumber); ok {
		o.SetHashedPhoneNumber(*str)
		return ok
	}
	return false
}

func (o *CapiEvent) SetIpAddress(unhashedIpAddress string) bool {
	if str, ok := NormAndHashStr(unhashedIpAddress); ok {
		o.SetHashedIpAddress(*str)
	}
	return false
}

func (o *CapiEvent) SetFirstName(unhashedFirstName string) bool {
	if str, ok := NormAndHashStr(unhashedFirstName); ok {
		o.SetHashedFirstNameSha(*str)
	} else {
		return false
	}
	if str, ok := NormAndSoundexStr(unhashedFirstName); ok {
		o.SetHashedFirstNameSdx(*str)
	}
	return false
}

func (o *CapiEvent) SetMiddleName(unhashedMiddleName string) bool {
	if str, ok := NormAndHashStr(unhashedMiddleName); ok {
		o.SetHashedMiddleNameSha(*str)
	} else {
		return false
	}
	if str, ok := NormAndSoundexStr(unhashedMiddleName); ok {
		o.SetHashedMiddleNameSdx(*str)
	}
	return false
}

func (o *CapiEvent) SetLastName(unhashedLastName string) bool {
	if str, ok := NormAndHashStr(unhashedLastName); ok {
		o.SetHashedLastNameSha(*str)
	} else {
		return false
	}
	if str, ok := NormAndSoundexStr(unhashedLastName); ok {
		o.SetHashedLastNameSdx(*str)
	}
	return false
}

func (o *CapiEvent) SetCity(unhashedCity string) bool {
	if str, ok := NormAndHashStr(unhashedCity); ok {
		o.SetHashedCitySha(*str)
	} else {
		return false
	}
	if str, ok := NormAndSoundexStr(unhashedCity); ok {
		o.SetHashedCitySdx(*str)
	}
	return false
}

func (o *CapiEvent) SetState(unhashedState string) bool {
	if str, ok := NormAndHashStr(unhashedState); ok {
		o.SetHashedStateSha(*str)
	} else {
		return false
	}
	if str, ok := NormAndSoundexStr(unhashedState); ok {
		o.SetHashedStateSdx(*str)
	}
	return false
}

func (o *CapiEvent) SetZip(unhashedZip string) bool {
	if str, ok := NormAndHashStr(unhashedZip); ok {
		o.SetHashedZip(*str)
		return ok
	}
	return false
}

func (o *CapiEvent) SetDob(unhashedDobStr string) bool {
	// TODO Parse and hash DD and MM after the format of unhashedDobStr is finalized
	return false
}
