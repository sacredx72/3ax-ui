func GenerateAWGClientConfig(server *model.AWGServer, client *model.AWGClient) string {
    var config strings.Builder
    
    config.WriteString("[Interface]\n")
    config.WriteString(fmt.Sprintf("PrivateKey = %s\n", client.PrivateKey))
    config.WriteString(fmt.Sprintf("Address = %s\n", client.Address))
    config.WriteString(fmt.Sprintf("DNS = %s\n", server.DNS))
    
    // AmneziaWG 2.0 параметры
    if server.S3 > 0 {
        config.WriteString(fmt.Sprintf("S3 = %d\n", server.S3))
    }
    if server.S4 > 0 {
        config.WriteString(fmt.Sprintf("S4 = %d\n", server.S4))
    }
    
    // Диапазоны заголовков (формат "min-max")
    if server.H1Max > server.H1Min {
        config.WriteString(fmt.Sprintf("H1 = %d-%d\n", server.H1Min, server.H1Max))
    } else if server.H1Min > 0 {
        config.WriteString(fmt.Sprintf("H1 = %d\n", server.H1Min))
    }
    // ... аналогично для H2, H3, H4 ...
    
    // CPS сигнатуры (только если заданы)
    if server.I1 != "" {
        config.WriteString(fmt.Sprintf("I1 = %s\n", server.I1))
    }
    if server.I2 != "" {
        config.WriteString(fmt.Sprintf("I2 = %s\n", server.I2))
    }
    if server.I3 != "" {
        config.WriteString(fmt.Sprintf("I3 = %s\n", server.I3))
    }
    if server.I4 != "" {
        config.WriteString(fmt.Sprintf("I4 = %s\n", server.I4))
    }
    if server.I5 != "" {
        config.WriteString(fmt.Sprintf("I5 = %s\n", server.I5))
    }
    
    config.WriteString("\n[Peer]\n")
    config.WriteString(fmt.Sprintf("PublicKey = %s\n", server.PublicKey))
    if server.PresharedKey != "" {
        config.WriteString(fmt.Sprintf("PresharedKey = %s\n", server.PresharedKey))
    }
    config.WriteString(fmt.Sprintf("Endpoint = %s:%d\n", server.Endpoint, server.Port))
    config.WriteString("AllowedIPs = 0.0.0.0/0, ::/0\n")
    
    return config.String()
}
