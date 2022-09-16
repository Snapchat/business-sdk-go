package businesssdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizePhoneNum(t *testing.T) {
	assert.Equal(t, "444472342344", normalizePhoneNum("+44-447-234-2344"))
	assert.Equal(t, "444472342344", normalizePhoneNum("0044-447-234-2344"))
	assert.Equal(t, "444472342344", normalizePhoneNum("+44-0447-234-2344"))
	assert.Equal(t, "444472342344", normalizePhoneNum("+44 0447 234 2344"))
	assert.Equal(t, "14472342344", normalizePhoneNum("447-234-2344"))
	assert.Equal(t, "14472342344", normalizePhoneNum("4472342344"))
	assert.Equal(t, "14472342344", normalizePhoneNum("+1-447-234-2344"))
	assert.Equal(t, "14472342344", normalizePhoneNum("(447)-234-2344"))
	assert.Equal(t, "14472342344", normalizePhoneNum("447 234-2344"))
	assert.Equal(t, "14472342344", normalizePhoneNum("+1-(447) 234 2344"))
}

func TestNormAndHashPhoneNum(t *testing.T) {
	var hashNum *string
	var ok bool

	// empty string
	hashNum, ok = NormAndHashPhoneNum("")
	assert.False(t, ok)
	assert.Nil(t, hashNum)

	testNum := "(447)-234-2344"
	hashNum, ok = NormAndHashPhoneNum(testNum)
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, Sha256("14472342344"), *hashNum)

	hashNum, ok = NormAndHashPhoneNum("+44-447-234-2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "19f25a18866af88fc60189c3af85010a0b971cf976053787fe21e004538c20a0", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("0044-447-234-2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "19f25a18866af88fc60189c3af85010a0b971cf976053787fe21e004538c20a0", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("+44-0447-234-2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "19f25a18866af88fc60189c3af85010a0b971cf976053787fe21e004538c20a0", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("+44 0447 234 2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "19f25a18866af88fc60189c3af85010a0b971cf976053787fe21e004538c20a0", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("447-234-2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "34053b331d29734f081bf74160cd33bc2f621b3b6f8ec991500bebf49f14e014", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("4472342344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "34053b331d29734f081bf74160cd33bc2f621b3b6f8ec991500bebf49f14e014", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("+1-447-234-2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "34053b331d29734f081bf74160cd33bc2f621b3b6f8ec991500bebf49f14e014", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("(447)-234-2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "34053b331d29734f081bf74160cd33bc2f621b3b6f8ec991500bebf49f14e014", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("447 234-2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "34053b331d29734f081bf74160cd33bc2f621b3b6f8ec991500bebf49f14e014", *hashNum)

	hashNum, ok = NormAndHashPhoneNum("+1-(447) 234 2344")
	assert.NotNil(t, hashNum)
	assert.True(t, ok)
	assert.Equal(t, "34053b331d29734f081bf74160cd33bc2f621b3b6f8ec991500bebf49f14e014", *hashNum)
}

func TestSoundex(t *testing.T) {
	assert.Equal(t, "H464", Soundex("Hello world"))
	assert.Equal(t, "T230", Soundex("test"))
	assert.Equal(t, "B252", Soundex("business sdk"))
	assert.Equal(t, "C516", Soundex("conversion api"))
	assert.Equal(t, "T235", Soundex("testing"))
}

func TestNormAndHashStr(t *testing.T) {
	var hashStr *string
	var ok bool

	// empty string
	hashStr, ok = NormAndHashStr("")
	assert.False(t, ok)
	assert.Nil(t, hashStr)

	// string with spaces only
	hashStr, ok = NormAndHashStr(" ")
	assert.False(t, ok)
	assert.Nil(t, hashStr)

	// test valid lowercase string
	testLower := "abc123"
	hashStr, ok = NormAndHashStr(testLower)
	assert.NotNil(t, hashStr)
	assert.True(t, ok)
	assert.Equal(t, Sha256(testLower), *hashStr)

	// test valid uppercase string
	testUpper := "ABC123"
	hashStr, ok = NormAndHashStr(testUpper)
	assert.NotNil(t, hashStr)
	assert.True(t, ok)
	assert.Equal(t, Sha256(testLower), *hashStr)

	// test mix case string
	testMix := "aBc123"
	hashStr, ok = NormAndHashStr(testMix)
	assert.NotNil(t, hashStr)
	assert.True(t, ok)
	assert.Equal(t, Sha256(testLower), *hashStr)

	hashStr, ok = NormAndHashStr("    test123@example.com  ")
	assert.NotNil(t, hashStr)
	assert.True(t, ok)
	assert.Equal(t, "20b82517b67108ed3a171539dc5dd1da605660249b02f38dd70f572a545c7d18", *hashStr)

	hashStr, ok = NormAndHashStr("   0.0.0.0  ")
	assert.NotNil(t, hashStr)
	assert.True(t, ok)
	assert.Equal(t, "19e36255972107d42b8cecb77ef5622e842e8a50778a6ed8dd1ce94732daca9e", *hashStr)

}
