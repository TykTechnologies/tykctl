package cloudcmd

/*func TestNewDeleteEnvCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockCloudClient(ctrl)
	prompt := mock.NewMockCloudPrompt(ctrl)
	config := mock.NewMockUserConfig(ctrl)
	factory := internal.CloudFactory{
		Client: m,
		Prompt: prompt,
		Config: config,
	}

	cmd := NewDeleteEnvCmd(factory)

	err := cmd.Execute()
	t.Log(err)
	///assert.Nil(t, t, err)
}*/
